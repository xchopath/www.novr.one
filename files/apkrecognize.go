package main

import (
	"fmt"
	"archive/zip"
	"os"
	"path/filepath"
)

type CommonApk struct {
	Tag string
	Files []string 
}

var ApkDictionary = []CommonApk{
	{
		Tag: "Flutter",
		Files: []string {
			"lib/armeabi-v7a/libflutter.so",
			"lib/arm64-v8a/libflutter.so",
			"lib/x86/libflutter.so",
			"lib/x86_64/libflutter.so",
		},
	},
	{
		Tag: "React Native",
		Files: []string {
			"lib/armeabi-v7a/libreactnativejni.so",
			"lib/arm64-v8a/libreactnativejni.so",
			"lib/x86/libreactnativejni.so",
			"lib/x86_64/libreactnativejni.so",
			"assets/index.android.bundle",
		},
	},
	{
		Tag: "Cordova / PhoneGap",
		Files: []string {
			"assets/www/index.html",
			"assets/www/cordova.js",
			"assets/www/cordova_plugins.js",
		},
	},
	{
		Tag: "Xamarin",
		Files: []string {
			"lib/armeabi-v7a/libmonodroid.so",
			"lib/arm64-v8a/libmonodroid.so",
			"lib/x86/libmonodroid.so",
			"lib/x86_64/libmonodroid.so",
			"lib/armeabi-v7a/libmonosgen-2.0.so",
			"lib/arm64-v8a/libmonosgen-2.0.so",
			"lib/x86/libmonosgen-2.0.so",
			"lib/x86_64/libmonosgen-2.0.so",
			"assemblies/Sikur.Monodroid.dll",
			"assemblies/Sikur.dll",
			"assemblies/Xamarin.Mobile.dll",
			"assemblies/mscorlib.dll",
		},
	},
	{
		Tag: "Corona SDK",
		Files: []string {
			"lib/armeabi-v7a/libcorona.so",
			"lib/arm64-v8a/libcorona.so",
			"lib/x86/libcorona.so",
			"lib/x86_64/libcorona.so",
			"assets/resource.car",
		},
	},
}

var banner = `
       .o.       ooooooooo.   oooo    oooo 
      .888.      '888   'Y88. '888   .8P'  
     .8'888.      888   .d88'  888  d8'    
    .8' '888.     888ooo88P'   88888[      
   .88ooo8888.    888          888'88b.    
  .8'     '888.   888          888  '88b.  
 o88o     o8888o o888o        o888o  o888o 
 ------ BASIC FRAMEWORK RECOGNIZER. ------
`

func main() {
	fmt.Println(banner)

	// Check command arguments
	args := os.Args
	if len(args) < 2 {
		cmdGo := filepath.Base(args[0])
		fmt.Println("ERROR: You did not provide APK file:", cmdGo, "~/path-to/example.apk")
		return
	}

	// Read APK
	ApkFile := args[1]
	ReadApk, error := zip.OpenReader(ApkFile)
	if error != nil {
		fmt.Println("ERROR: There was an error:", error)
		return
	}
	defer ReadApk.Close()
	fmt.Println("INFO: Checking", ApkFile + "...")

	// Matching common paths
	Found := 0
	for _, file := range ReadApk.File {
		for _, ApkCommon := range ApkDictionary {
			for _, ApkCommon_File := range ApkCommon.Files {
				if file.Name == ApkCommon_File {
					Found++
					fmt.Println("INFO: Found \"" + file.Name + "\" file, It should be [" + ApkCommon.Tag + "].")
				}
			}
		}
	}

	// If no file is found then it is considered as native
	if Found == 0 {
		fmt.Println("INFO: Could not recognize \"" + ApkFile + "\", perhaps It is [Native].")
	}
	return 
}
