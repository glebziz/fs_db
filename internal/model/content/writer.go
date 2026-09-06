package content

import (
	"errors"
	"io"
	"sync"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/utils/async"
)

type content struct {
	sync.WaitGroup

	rw  io.ReadWriteCloser
	pos model.Position
	ch  chan model.Content

	errM sync.RWMutex
	err  error
}

func New() (*content, model.Contents) {
	c := content{
		ch: make(chan model.Content, 1),
	}

	return &c, c.ch
}

func (c *content) Close() error {
	close(c.ch)

	if c.rw != nil {
		c.rw.Close()
	}

	c.Wait()

	return c.err
}

func (c *content) Seek(offset int64, whence int) (int64, error) {
	err := c.checkErr()
	if err != nil {
		return 0, err
	}

	seek := model.NewSeek(offset, whence)
	if !c.pos.NeedSeek(seek) {
		return c.pos.Pos(), nil
	}

	if c.rw != nil {
		c.rw.Close()
		c.rw = nil
	}

	err = c.pos.Seek(seek)
	if err != nil {
		return 0, err
	}

	return c.pos.Pos(), nil
}

func (c *content) Write(p []byte) (n int, err error) {
	err = c.checkErr()
	if err != nil {
		return 0, err
	}

	if len(p) == 0 {
		return 0, nil
	}

	defer func() {
		_ = c.pos.Seek(model.Seek{
			Pos: int64(len(p)),
			Dir: model.SeekCurrent,
		})
	}()

	if c.rw == nil {
		c.rw = async.NewReadWriter()
		c.ch <- model.Content{
			Reader: c.rw,
			Seek: model.Seek{
				Pos: c.pos.Pos(),
				Dir: model.SeekStart,
			},
		}
	}

	return c.rw.Write(p)
}

func (c *content) WriteAt(p []byte, off int64) (n int, err error) {
	err = c.checkErr()
	if err != nil {
		return 0, err
	}

	_, err = c.Seek(off, io.SeekStart)
	if err != nil {
		return 0, err
	}

	return c.Write(p)
}

func (c *content) SetError(err error) {
	c.errM.Lock()
	defer c.errM.Unlock()

	c.err = errors.Join(c.err, err)
}

func (c *content) checkErr() error {
	c.errM.RLock()
	defer c.errM.RUnlock()

	return c.err
}
