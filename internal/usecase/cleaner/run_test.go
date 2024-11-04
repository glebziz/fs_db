package cleaner

//func TestCleaner_Run(t *testing.T) {
//	for _, tc := range []struct {
//		name string
//		err  error
//	}{
//		{
//			name: "success",
//		},
//		{
//			name: "schedule error",
//			err:  assert.AnError,
//		},
//	} {
//		tc := tc
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//
//			td := newTestDeps(t)
//
//			td.txRepo.EXPECT().
//				Oldest(gomock.Any()).
//				Return(nil, assert.AnError)
//
//			td.sched.EXPECT().
//				ScheduleJob(gomock.Any(), gomock.Any()).
//				DoAndReturn(func(detail *quartz.JobDetail, _ quartz.Trigger) error {
//					err := detail.Job().Execute(testCtx)
//					require.NoError(t, err)
//
//					return tc.err
//				})
//
//			if tc.err == nil {
//				td.sched.EXPECT().
//					Start(gomock.Any())
//			}
//
//			cl := td.newCleaner()
//
//			err := cl.Run(testCtx)
//			if tc.err != nil {
//				require.ErrorIs(t, err, tc.err)
//			} else {
//				require.NoError(t, err)
//			}
//		})
//	}
//}
