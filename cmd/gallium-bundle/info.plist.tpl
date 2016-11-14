<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>BuildMachineOSBuild</key>
	<string>14F1509</string>
	<key>CFBundleIdentifier</key>
	<string>{{.BundleIdentifier}}</string>
	<key>CFBundleName</key>
	<string>{{.BundleName}}</string>
	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>MacOSX</string>
	</array>
	<key>DTCompiler</key>
	<string>com.apple.compilers.llvm.clang.1_0</string>
	<key>DTPlatformBuild</key>
	<string>7C1002</string>
	<key>DTPlatformVersion</key>
	<string>GM</string>
	<key>DTSDKBuild</key>
	<string>15C43</string>
	<key>DTSDKName</key>
	<string>macosx10.11</string>
	<key>DTXcode</key>
	<string>0721</string>
	<key>DTXcodeBuild</key>
	<string>7C1002</string>
	<key>NSSupportsAutomaticGraphicsSwitching</key>
	<true/>
	{{range $key, $val := .Extras}}
		<key>{{$key}}</key>
		<string>{{$val}}</string>
	{{end}}
</dict>
</plist>
