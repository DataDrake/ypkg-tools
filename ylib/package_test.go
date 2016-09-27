package ylib

import (
	"testing"
	"strings"
	"bytes"
)

var packageYaml = `name	   : nano
version    : 2.7.0
release    : 57
source     :
    - https://www.nano-editor.org/dist/v2.7/nano-2.7.0.tar.xz : f86af39514ae74e20bef3c29cd01d1090a9aca772a70e9c9f9e70c3d14b39521
license    :
    - GPL-3.0
summary    : Small, friendly text editor inspired by Pico
component  : editor
description: |
    GNU nano is an easy-to-use text editor originally designed as a replacement for Pico, the ncurses-based editor from the non-free mailer package P$
setup	   : |
    %patch -p1 < $pkgfiles/0001-Use-a-stateless-configuration.patch
    autoreconf -vdi
    %configure --enable-utf8 --docdir=/usr/share/doc/nano
build	   : |
    %make
install    : |
    %make_install
    install -D -m 00644 $pkgfiles/nanorc $installdir/usr/share/defaults/nano/nanorc
    install -D -m 00644 $pkgfiles/git.nanorc $installdir/usr/share/nano/git.nanorc
`

func TestPackageYML_Read(t *testing.T) {
	raw := strings.NewReader(packageYaml)
	pkg := &PackageYML{}
	err := pkg.Read(raw)
	if err != nil {
		t.Errorf("Failed to parse, reason: %s",err.Error())
	}
	if pkg.Name != "nano" {
		t.Errorf("Should be '%s', found: '%s'","nano",pkg.Name)
	}
	if pkg.Version != "2.7.0" {
		t.Errorf("Should be '%s', found: '%s'","2.7.0",pkg.Version)
	}
	if len(pkg.Source) != 1 {
		t.Errorf("Should be '%d', found: '%d'",1,len(pkg.Source))
	}
}

func TestPackageYML_Write(t *testing.T) {
	pkg := &PackageYML{}
	src := make(map[string]string)
	src["http://example.com/file-0.6.4.tar.xz"] = "22222"
	pkg.Source = make([]map[string]string,0)
	pkg.Source = append(pkg.Source,src)
	out := bytes.NewBuffer(make([]byte,0))
	_, err := pkg.Write(out)
	t.Log(out.String())
	if err == nil {
		t.Error("Error should have occured")
	}
}