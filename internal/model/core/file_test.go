package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db/internal/model"
	"github.com/glebziz/fs_db/internal/model/sequence"
)

func TestFile_Lock(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    *file
	}{
		{
			name: "success",
			f:    &file{},
		},
		{
			name: "success with nil file",
			f:    nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.f.Lock()
		})
	}
}

func TestFile_RLock(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    *file
	}{
		{
			name: "success",
			f:    &file{},
		},
		{
			name: "success with nil file",
			f:    nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.f.RLock()
		})
	}
}

func TestFile_Unlock(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    *file
	}{
		{
			name: "success",
			f: func() *file {
				f := &file{}
				f.Lock()
				return f
			}(),
		},
		{
			name: "success with nil file",
			f: func() *file {
				return nil
			}(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.f.Unlock()
		})
	}
}

func TestFile_RUnlock(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    *file
	}{
		{
			name: "success",
			f: func() *file {
				f := &file{}
				f.RLock()
				return f
			}(),
		},
		{
			name: "success with nil file",
			f: func() *file {
				return nil
			}(),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.f.RUnlock()
		})
	}
}

func TestFile_PushBack(t *testing.T) {
	for _, tc := range []struct {
		name     string
		f        func() *file
		n        *Node[model.File]
		requireF func(t *testing.T, f *file)
	}{
		{
			name: "push to nil file",
			f: func() *file {
				return nil
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey,
					ContentId: testContentId,
					Seq:       testSeq,
				},
			},
			requireF: func(t *testing.T, f *file) {
				require.Nil(t, f)
			},
		},
		{
			name: "push nil node",
			f: func() *file {
				return &file{}
			},
			n: nil,
			requireF: func(t *testing.T, f *file) {
				require.Empty(t, f.arr)
				require.True(t, f.l.IsEmpty())
			},
		},
		{
			name: "push to empty file",
			f: func() *file {
				return &file{}
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey,
					ContentId: testContentId,
					Seq:       testSeq,
				},
			},
			requireF: func(t *testing.T, f *file) {
				require.Len(t, f.arr, 1)
				require.True(t, f.arr[0] == f.l.Back())
			},
		},
		{
			name: "push to non empty file",
			f: func() *file {
				f := &file{}
				n := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}

				f.l.PushBack(n)
				f.arr = append(f.arr, n)

				return f
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey2,
					ContentId: testContentId2,
					Seq:       testSeq + 1,
				},
			},
			requireF: func(t *testing.T, f *file) {
				require.Len(t, f.arr, 2)
				require.True(t, f.arr[1] == f.l.Back())
				require.True(t, f.arr[0] == f.l.Front())
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := tc.f()
			f.PushBack(tc.n)
			tc.requireF(t, f)
		})
	}
}

func TestFile_PopBack(t *testing.T) {
	for _, tc := range []struct {
		name     string
		f        func() *file
		requireF func(t *testing.T, f *file)
		n        *Node[model.File]
	}{
		{
			name: "pop from nil file",
			f: func() *file {
				return nil
			},
			requireF: func(t *testing.T, f *file) {
				require.Nil(t, f)
			},
			n: nil,
		},
		{
			name: "pop from empty file",
			f: func() *file {
				return &file{}
			},
			requireF: func(t *testing.T, f *file) {
				require.Empty(t, f.arr)
				require.True(t, f.l.IsEmpty())
			},
			n: nil,
		},
		{
			name: "pop from non empty file",
			f: func() *file {
				f := &file{}

				n := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}

				f.l.PushBack(n)
				f.arr = append(f.arr, n)

				return f
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey,
					ContentId: testContentId,
					Seq:       testSeq,
				},
			},
			requireF: func(t *testing.T, f *file) {
				require.Empty(t, f.arr)
				require.True(t, f.l.IsEmpty())
			},
		},
		{
			name: "pop from file with multiple nodes",
			f: func() *file {
				f := &file{}
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 1,
					},
				}

				f.l.PushBack(n1)
				f.l.PushBack(n2)
				f.arr = append(f.arr, n1, n2)

				return f
			},
			requireF: func(t *testing.T, f *file) {
				require.Len(t, f.arr, 1)
				require.True(t, f.arr[0] == f.l.Back())
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey2,
					ContentId: testContentId2,
					Seq:       testSeq + 1,
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := tc.f()
			n := f.PopBack()
			tc.requireF(t, f)
			require.Equal(t, tc.n, n)
		})
	}
}

func TestFile_PopFront(t *testing.T) {
	for _, tc := range []struct {
		name     string
		f        func() *file
		requireF func(t *testing.T, f *file)
		n        *Node[model.File]
	}{
		{
			name: "pop from nil file",
			f: func() *file {
				return nil
			},
			requireF: func(t *testing.T, f *file) {
				require.Nil(t, f)
			},
			n: nil,
		},
		{
			name: "pop from empty file",
			f: func() *file {
				return &file{}
			},
			requireF: func(t *testing.T, f *file) {
				require.Empty(t, f.arr)
				require.True(t, f.l.IsEmpty())
			},
			n: nil,
		},
		{
			name: "pop from non empty file",
			f: func() *file {
				f := &file{}
				n := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}

				f.l.PushBack(n)
				f.arr = append(f.arr, n)

				return f
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey,
					ContentId: testContentId,
					Seq:       testSeq,
				},
			},
			requireF: func(t *testing.T, f *file) {
				require.Empty(t, f.arr)
				require.True(t, f.l.IsEmpty())
			},
		},
		{
			name: "pop from file with multiple nodes",
			f: func() *file {
				f := &file{}
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 1,
					},
				}

				f.l.PushBack(n1)
				f.l.PushBack(n2)
				f.arr = append(f.arr, n1, n2)

				return f
			},
			requireF: func(t *testing.T, f *file) {
				require.Len(t, f.arr, 1)
				require.True(t, f.arr[0] == f.l.Back())
			},
			n: &Node[model.File]{
				v: model.File{
					Key:       testKey,
					ContentId: testContentId,
					Seq:       testSeq,
				},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f := tc.f()
			n := f.PopFront()
			tc.requireF(t, f)
			require.Equal(t, tc.n, n)
		})
	}
}

func TestFile_Latest(t *testing.T) {
	for _, tc := range []struct {
		name string
		f    func() *file
		file model.File
	}{
		{
			name: "latest from nil file",
			f: func() *file {
				return nil
			},
			file: model.File{},
		},
		{
			name: "latest from empty file",
			f: func() *file {
				return &file{}
			},
			file: model.File{},
		},
		{
			name: "latest from non empty file",
			f: func() *file {
				n := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}

				f := &file{}
				f.PushBack(n)
				return f
			},
			file: model.File{
				Key:       testKey,
				ContentId: testContentId,
				Seq:       testSeq,
			},
		},
		{
			name: "latest from file with multiple nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 1,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			file: model.File{
				Key:       testKey2,
				ContentId: testContentId2,
				Seq:       testSeq + 1,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			file := tc.f().Latest()
			require.Equal(t, tc.file, file)
		})
	}
}

func TestFile_LastBefore(t *testing.T) {
	for _, tc := range []struct {
		name      string
		f         func() *file
		beforeSeq sequence.Seq
		file      model.File
	}{
		{
			name: "last before from nil file",
			f: func() *file {
				return nil
			},
			beforeSeq: testSeq,
			file:      model.File{},
		},
		{
			name: "last before from empty file",
			f: func() *file {
				return &file{}
			},
			beforeSeq: testSeq,
			file:      model.File{},
		},
		{
			name: "last before from file with multiple after nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq + 1,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 2,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			beforeSeq: testSeq,
			file:      model.File{},
		},
		{
			name: "last before from file with multiple before nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq - 2,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq - 1,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			beforeSeq: testSeq,
			file: model.File{
				Key:       testKey2,
				ContentId: testContentId2,
				Seq:       testSeq - 1,
			},
		},
		{
			name: "last before from file with multiple nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq - 1,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 1,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			beforeSeq: testSeq,
			file: model.File{
				Key:       testKey,
				ContentId: testContentId,
				Seq:       testSeq - 1,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			file := tc.f().LastBefore(tc.beforeSeq)
			require.Equal(t, tc.file, file)
		})
	}
}

func TestFile_IterateBeforeTs(t *testing.T) {
	for _, tc := range []struct {
		name      string
		f         func() *file
		beforeSeq sequence.Seq
		files     []model.File
	}{
		{
			name: "iterate nil file",
			f: func() *file {
				return nil
			},
			beforeSeq: testSeq,
			files:     []model.File{},
		},
		{
			name: "iterate empty file",
			f: func() *file {
				return &file{}
			},
			beforeSeq: testSeq,
			files:     []model.File{},
		},
		{
			name: "iterate file with one before node",
			f: func() *file {
				n := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq - 1,
					},
				}

				f := &file{}
				f.PushBack(n)
				return f
			},
			beforeSeq: testSeq,
			files:     []model.File{},
		},
		{
			name: "iterate file with one after node",
			f: func() *file {
				n := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq + 1,
					},
				}

				f := &file{}
				f.PushBack(n)
				return f
			},
			beforeSeq: testSeq,
			files:     []model.File{},
		},
		{
			name: "iterate file with multiple after nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq + 1,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 2,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			beforeSeq: testSeq,
			files:     []model.File{},
		},
		{
			name: "iterate file with multiple before nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq - 2,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq - 1,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			beforeSeq: testSeq,
			files: []model.File{{
				Key:       testKey,
				ContentId: testContentId,
				Seq:       testSeq - 2,
			}},
		},
		{
			name: "iterate file with multiple nodes",
			f: func() *file {
				n1 := &Node[model.File]{
					v: model.File{
						Key:       testKey,
						ContentId: testContentId,
						Seq:       testSeq - 1,
					},
				}
				n2 := &Node[model.File]{
					v: model.File{
						Key:       testKey2,
						ContentId: testContentId2,
						Seq:       testSeq + 1,
					},
				}

				f := &file{}
				f.PushBack(n1)
				f.PushBack(n2)
				return f
			},
			beforeSeq: testSeq,
			files:     []model.File{},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var (
				f     = tc.f()
				files = make([]model.File, 0, len(tc.files))
			)
			for v := range f.IterateBeforeSeq(tc.beforeSeq) {
				files = append(files, v)
			}
			for v := range f.IterateBeforeSeq(tc.beforeSeq) {
				_ = v
				break
			}

			require.Equal(t, len(tc.files), len(files))
			require.Equal(t, tc.files, files)
		})
	}
}

func Test_binarySearch(t *testing.T) {
	for _, tc := range []struct {
		name string
		arr  []*Node[model.File]
		seq  sequence.Seq
		n    *Node[model.File]
	}{
		{
			name: "empty",
			arr:  nil,
			seq:  testSeq,
			n:    nil,
		},
		{
			name: "single equal",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq,
				},
			}},
			seq: testSeq,
			n:   nil,
		},
		{
			name: "single after",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq + 1,
				},
			}},
			seq: testSeq,
			n:   nil,
		},
		{
			name: "single before",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq - 1,
				},
			}},
			seq: testSeq,
			n: &Node[model.File]{
				v: model.File{
					Seq: testSeq - 1,
				},
			},
		},
		{
			name: "multiple before",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq - 10,
				},
			}, {
				v: model.File{
					Seq: testSeq - 5,
				},
			}, {
				v: model.File{
					Seq: testSeq - 1,
				},
			}},
			seq: testSeq,
			n: &Node[model.File]{
				v: model.File{
					Seq: testSeq - 1,
				},
			},
		},
		{
			name: "multiple before and equal",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq - 10,
				},
			}, {
				v: model.File{
					Seq: testSeq - 5,
				},
			}, {
				v: model.File{
					Seq: testSeq - 1,
				},
			}, {
				v: model.File{
					Seq: testSeq,
				},
			}},
			seq: testSeq,
			n: &Node[model.File]{
				v: model.File{
					Seq: testSeq - 1,
				},
			},
		},
		{
			name: "multiple after",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq + 1,
				},
			}, {
				v: model.File{
					Seq: testSeq + 5,
				},
			}, {
				v: model.File{
					Seq: testSeq + 10,
				},
			}},
			seq: testSeq,
			n:   nil,
		},
		{
			name: "multiple after and equal",
			arr: []*Node[model.File]{{
				v: model.File{
					Seq: testSeq,
				},
			}, {
				v: model.File{
					Seq: testSeq + 1,
				},
			}, {
				v: model.File{
					Seq: testSeq + 5,
				},
			}, {
				v: model.File{
					Seq: testSeq + 10,
				},
			}},
			seq: testSeq,
			n:   nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			n := binarySearch(tc.arr, tc.seq)
			require.Equal(t, tc.n, n)
		})
	}
}
