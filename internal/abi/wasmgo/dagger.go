package wasmgo

import "log"

func (w *wasmGo) daggerOpenFD(sp int32) {
	fname := string(w.goLoadSlice(sp + 8))

	log.Printf("dagger.OpenFD(%q)", fname)

	w.setInt64(sp+8, w.child.OpenFD(fname, 0))
}
