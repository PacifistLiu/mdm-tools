package main

func getApp(requestedApp string) string {
	apps := make(map[string]string)
	apps["adb"] = `{"pkg": "com.example.adb"}`
	apps["ff"] = `{"pkg": "org.mozilla.firefox"}`
	apps["set"] = `{"pkg": "com.android.settings"}`

	return apps[requestedApp]
}
