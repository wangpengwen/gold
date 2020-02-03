package layers

import (
	"fmt"

	g "gorgonia.org/gorgonia"
)

// Activation is an activation function.
type Activation interface {
	// Fwd is a foward pass through x.
	Fwd(x *g.Node) (*g.Node, error)

	// Clone the activation.
	Clone() Activation
}

// SigmoidActivation is a sigmoid activation layer.
type SigmoidActivation struct{}

// Sigmoid activation function.
var Sigmoid = &SigmoidActivation{}

// NewSigmoid returns a new sigmoid activation layer.
func NewSigmoid() *SigmoidActivation {
	return &SigmoidActivation{}
}

// Fwd is a foward pass through the layer.
func (s *SigmoidActivation) Fwd(x *g.Node) (*g.Node, error) {
	return g.Sigmoid(x)
}

// Learnables returns all learnable nodes within this layer.
func (s *SigmoidActivation) Learnables() (n g.Nodes) {
	return n
}

// Compile the layer.
func (s *SigmoidActivation) Compile(x *g.Node, opts ...LayerOpt) {}

// Clone the activation.
func (s *SigmoidActivation) Clone() Activation {
	return NewSigmoid()
}

// TanhActivation is a tanh activation layer.
type TanhActivation struct{}

// Tanh activation.
var Tanh = &TanhActivation{}

// NewTanh returns a new tanh activation layer.
func NewTanh() *TanhActivation {
	return &TanhActivation{}
}

// Fwd is a foward pass through the layer.
func (t *TanhActivation) Fwd(x *g.Node) (*g.Node, error) {
	return g.Tanh(x)
}

// Learnables returns all learnable nodes within this layer.
func (t *TanhActivation) Learnables() (n g.Nodes) {
	return n
}

// Compile the layer.
func (t *TanhActivation) Compile(x *g.Node, opts ...LayerOpt) {}

// Clone the activation.
func (t *TanhActivation) Clone() Activation {
	return NewTanh()
}

// ReLUActivation is a relu activation layer.
type ReLUActivation struct{}

// ReLU activation.
var ReLU = &ReLUActivation{}

// NewReLU returns a new relu activation layer.
func NewReLU() *ReLUActivation {
	return &ReLUActivation{}
}

// Fwd is a foward pass through the layer.
func (r *ReLUActivation) Fwd(x *g.Node) (*g.Node, error) {
	return g.Rectify(x)
}

// Learnables returns all learnable nodes within this layer.
func (r *ReLUActivation) Learnables() (n g.Nodes) {
	return n
}

// Compile the layer.
func (r *ReLUActivation) Compile(x *g.Node, opts ...LayerOpt) {}

// Clone the activation.
func (r *ReLUActivation) Clone() Activation {
	return NewReLU()
}

// LeakyReLUActivation is a leaky relu activation layer.
type LeakyReLUActivation struct {
	alpha float64
}

// LeakyReLU is default leaky relu activation.
var LeakyReLU = &LeakyReLUActivation{0.01}

// NewLeakyReLU returns a new leaky relu activation layer.
func NewLeakyReLU(alpha float64) *LeakyReLUActivation {
	return &LeakyReLUActivation{alpha: alpha}
}

// Fwd is a foward pass through the layer.
func (r *LeakyReLUActivation) Fwd(x *g.Node) (*g.Node, error) {
	return g.LeakyRelu(x, r.alpha)
}

// Learnables returns all learnable nodes within this layer.
func (r *LeakyReLUActivation) Learnables() (n g.Nodes) {
	return n
}

// Compile the layer.
func (r *LeakyReLUActivation) Compile(x *g.Node, opts ...LayerOpt) {}

// Clone the activation.
func (r *LeakyReLUActivation) Clone() Activation {
	return NewLeakyReLU(r.alpha)
}

// SoftmaxActivation is a softmax activation layer.
type SoftmaxActivation struct {
	axis []int
}

// SoftMax is the default softmax activation.
var SoftMax = &SoftmaxActivation{}

// NewSoftmax returns a new leaky softmax activation layer.
func NewSoftmax(axis ...int) *SoftmaxActivation {
	// if len(axis) == 0 {
	// 	axis = append(axis, 0)
	// }
	return &SoftmaxActivation{axis: axis}
}

// Fwd is a foward pass through the layer.
func (s *SoftmaxActivation) Fwd(x *g.Node) (*g.Node, error) {
	fmt.Printf("running softmax with x shape: %v dims: %v \n", x.Shape(), x.Dims())
	return g.SoftMax(x, s.axis...)
}

// Learnables returns all learnable nodes within this layer.
func (s *SoftmaxActivation) Learnables() (n g.Nodes) {
	return n
}

// Compile the layer.
func (s *SoftmaxActivation) Compile(x *g.Node, opts ...LayerOpt) {}

// Clone the activation.
func (s *SoftmaxActivation) Clone() Activation {
	return NewSoftmax(s.axis...)
}

// LinearActivation is a linear (identity) activation layer.
type LinearActivation struct{}

// Linear activation.
var Linear = &LinearActivation{}

// NewLinear is a linear activation layer.
func NewLinear() *LinearActivation {
	return &LinearActivation{}
}

// Fwd is a foward pass through the layer.
func (l *LinearActivation) Fwd(x *g.Node) (*g.Node, error) {
	return x, nil
}

// Learnables returns all learnable nodes within this layer.
func (l *LinearActivation) Learnables() (n g.Nodes) {
	return n
}

// Compile the layer.
func (l *LinearActivation) Compile(x *g.Node, opts ...LayerOpt) {}

// Clone the activation.
func (l *LinearActivation) Clone() Activation {
	return NewLinear()
}
