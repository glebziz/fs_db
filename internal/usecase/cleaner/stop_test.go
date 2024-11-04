package cleaner

//func TestCleaner_Stop(t *testing.T) {
//	for _, tc := range []struct {
//		name string
//		err  error
//	}{
//		{
//			name: "success",
//		},
//		{
//			name: "clear error",
//			err:  assert.AnError,
//		},
//	} {
//		tc := tc
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//
//			td := newTestDeps(t)
//
//			td.sched.EXPECT().
//				Stop()
//
//			td.sched.EXPECT().
//				Wait(gomock.Any())
//
//			td.sched.EXPECT().
//				Clear().
//				Return(tc.err)
//
//			cl := td.newCleaner()
//
//			err := cl.Stop(testCtx)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
