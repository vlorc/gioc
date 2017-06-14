package gioc


import (
	"testing"
	"github.com/vlorc/gioc/types"
	"github.com/vlorc/gioc/binder"
	"github.com/vlorc/gioc/register"
)

func test_register(t *testing.T,r types.Register) {
	if nil == r {
		t.Errorf("can't allocate a Register")
	}

	bind := binder.NewNameBinder()
	err := r.RegisterBinder(bind,types.BinderFactoryType)
	if nil != err {
		t.Errorf("can't register a Binder error : %s",err.Error())
	}

	if bind != r.AsBinder(types.BinderFactoryType) {
		t.Errorf("can't matching Binder,were modified")
	}
}

func test_registerFactory(t *testing.T,f types.RegisterFactory) {
	if nil == f {
		t.Errorf("can't allocate a RegisterFactory")
	}

	r,err := f.Instance(binder.NewBinderFactory())
	if nil != err {
		t.Errorf("can't allocate a Register error : %s",err.Error())
	}
	test_register(t,r)
}

func Test_Register(t *testing.T) {

	test_register(t,register.NewRegister(binder.NewBinderFactory()))
}

func Test_RegisterFactory(t *testing.T) {
	test_registerFactory(t,register.NewRegisterFactory())
}

