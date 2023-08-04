package downloader

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ofdl/ofdl/model"
	"gorm.io/gorm"
)

type LocalDownloader struct {
	db   *gorm.DB
	root string
}

func NewLocalDownloader(db *gorm.DB, root string) (Downloader, error) {
	return &LocalDownloader{
		db:   db,
		root: root,
	}, nil
}

// DownloadMany implements Downloader.
func (d *LocalDownloader) DownloadMany(mm []model.DownloadableMedia) <-chan error {
	d1 := make(chan error)

	go func() {
		sem := make(chan struct{}, 5)
		for _, m := range mm {
			sem <- struct{}{}
			go func(m model.DownloadableMedia) {
				defer func() { <-sem }()

				p, done := d.DownloadOne(m)
				for {
					select {
					case <-p:
					case err := <-done:
						d1 <- err
						return
					}
				}
			}(m)
		}
		for i := 0; i < cap(sem); i++ {
			sem <- struct{}{}
		}
		close(d1)
	}()

	return d1
}

// DownloadOne implements Downloader.
func (d *LocalDownloader) DownloadOne(m model.DownloadableMedia) (<-chan float64, <-chan error) {
	progress := make(chan float64)
	done := make(chan error)

	go func() {
		defer close(progress)

		if m.URL() == "" {
			done <- m.MarkDownloaded(d.db)
			return
		}

		resp, err := http.Get(m.URL())
		if err != nil {
			done <- err
			return
		}
		defer resp.Body.Close()

		totalSize := resp.ContentLength
		progressReader := ProgressReader(resp.Body, totalSize, progress)

		if err := os.MkdirAll(filepath.Join(d.root, m.Directory()), 0755); err != nil {
			done <- err
			return
		}
		out, err := os.Create(filepath.Join(d.root, m.Directory(), m.Filename()))
		if err != nil {
			done <- err
			return
		}
		defer out.Close()

		_, err = io.Copy(out, progressReader)
		if err != nil {
			done <- err
			return
		}

		done <- m.MarkDownloaded(d.db)
	}()

	return progress, done
}

func ProgressReader(reader io.Reader, totalSize int64, progressChan chan<- float64) io.Reader {
	return &progressReaderStruct{
		reader:       reader,
		totalSize:    totalSize,
		readSoFar:    0,
		progressChan: progressChan,
	}
}

type progressReaderStruct struct {
	reader       io.Reader
	totalSize    int64
	readSoFar    int64
	progressChan chan<- float64
}

func (prs *progressReaderStruct) Read(p []byte) (int, error) {
	n, err := prs.reader.Read(p)
	prs.readSoFar += int64(n)
	progress := float64(prs.readSoFar) / float64(prs.totalSize) // * 100.0
	prs.progressChan <- progress
	return n, err
}
