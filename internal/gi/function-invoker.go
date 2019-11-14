package gi

// #include <girepository.h>
import "C"
import "C"

type FunctionInvoker struct {
	info *C.GIFunctionInfo
}

func FunctionInvokerNew(funcName string) *FunctionInvoker {
	return nil
}
