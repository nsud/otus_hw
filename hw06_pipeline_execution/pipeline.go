package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	var (
		countStg = len(stages)
		out      = make(Bi)
	)

	if countStg == 0 {
		close(out)
		return out
	}
	firstStg := stages[0](out)

	go func() {
		defer close(out)
		for {
			select {
			case val, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- val:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()

	if countStg > 1 {
		return ExecutePipeline(firstStg, done, stages[1:]...)
	}

	return firstStg
}
