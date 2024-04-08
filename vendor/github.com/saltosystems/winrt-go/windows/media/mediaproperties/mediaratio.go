// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package mediaproperties

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

const SignatureMediaRatio string = "rc(Windows.Media.MediaProperties.MediaRatio;{d2d0fee5-8929-401d-ac78-7d357e378163})"

type MediaRatio struct {
	ole.IUnknown
}

func (impl *MediaRatio) SetNumerator(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaRatio))
	defer itf.Release()
	v := (*iMediaRatio)(unsafe.Pointer(itf))
	return v.SetNumerator(value)
}

func (impl *MediaRatio) GetNumerator() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaRatio))
	defer itf.Release()
	v := (*iMediaRatio)(unsafe.Pointer(itf))
	return v.GetNumerator()
}

func (impl *MediaRatio) SetDenominator(value uint32) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaRatio))
	defer itf.Release()
	v := (*iMediaRatio)(unsafe.Pointer(itf))
	return v.SetDenominator(value)
}

func (impl *MediaRatio) GetDenominator() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiMediaRatio))
	defer itf.Release()
	v := (*iMediaRatio)(unsafe.Pointer(itf))
	return v.GetDenominator()
}

const GUIDiMediaRatio string = "d2d0fee5-8929-401d-ac78-7d357e378163"
const SignatureiMediaRatio string = "{d2d0fee5-8929-401d-ac78-7d357e378163}"

type iMediaRatio struct {
	ole.IInspectable
}

type iMediaRatioVtbl struct {
	ole.IInspectableVtbl

	SetNumerator   uintptr
	GetNumerator   uintptr
	SetDenominator uintptr
	GetDenominator uintptr
}

func (v *iMediaRatio) VTable() *iMediaRatioVtbl {
	return (*iMediaRatioVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iMediaRatio) SetNumerator(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetNumerator,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaRatio) GetNumerator() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetNumerator,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iMediaRatio) SetDenominator(value uint32) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetDenominator,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in uint32
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iMediaRatio) GetDenominator() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDenominator,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}