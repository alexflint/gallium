default: lib demo

force:

lib: force
	#cd lib; ./script/build || (cd -; exit 1); cd -
	PYTHONPATH=lib/vendor/gyp/pylib \
		python lib/vendor/gyp/gyp_main.py --depth lib lib/gallium.gyp || exit 1
	xcodebuild -project lib/gallium.xcodeproj || exit 1
	rm -rf dist/Gallium.framework
	mv lib/build/Debug/Gallium.framework dist/

demo: lib force
	go build ./cmd/demo || exit 1
	go install ./cmd/gallium-bundle || exit 1
	gallium-bundle demo

menudemo: lib force
	go build ./cmd/menudemo || exit 1
	go install ./cmd/gallium-bundle || exit 1
	gallium-bundle menudemo

run: demo
	open demo.app
