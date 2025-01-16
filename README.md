# Go HTTP Proxy Server with VPN Sharing

A highly efficient and configurable HTTP proxy server written in Go. This project is designed to enable VPN sharing across devices by acting as a proxy server. It supports dynamic port configuration, graceful shutdown, and status monitoring. Additionally, it includes an Android AAR library for easy integration into mobile applications.

---

## Features

- **HTTP Proxy Server**
- **VPN Sharing Support**
- **Dynamic Port Configuration**
- **Status Monitoring**
- **Graceful Start/Stop Mechanisms**
- **Android AAR for Mobile Integration**

---

## Getting Started

### Prerequisites

Ensure you have the following tools installed:

- **Go 1.23+** ([Download Here](https://go.dev/dl/))
- **Android SDK** (for building the AAR)
- **gomobile** command-line tool ([Installation Guide](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile))

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/khaledagn/go-http-proxy-server.git
   cd go-http-proxy-server
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Install `gomobile` and initialize it:
   ```bash
   go install golang.org/x/mobile/cmd/gomobile@latest
   gomobile init
   ```

---

## AAR Generation Guide

The project includes an Android AAR library to enable starting and stopping the proxy server directly from Android applications.

### Steps to Generate the AAR

1. Run the provided script to generate the AAR file:
   ```bash
   ./generate_aar.sh
   ```

2. The generated AAR will be located in the `output/` directory:
   ```
   output/proxyserver.aar
   ```

### Adding the AAR to Your Android Project

1. Copy `proxyserver.aar` to your Android project’s `libs` directory.
2. Update your app’s `build.gradle` (Module) file:
   ```gradle
   repositories {
       flatDir {
           dirs 'libs'
       }
   }

   dependencies {
       implementation(name: 'proxyserver', ext: 'aar')
   }
   ```
3. Sync the project in Android Studio.

### Using the AAR in Your Java Code

The AAR provides two primary methods for interacting with the proxy server:

- **`Proxyserver.StartProxy(String listenPort)`**: Starts the proxy server on the specified port.
- **`Proxyserver.StopProxy()`**: Stops the proxy server.

#### Java Code Example

```java
import proxyserver.Proxyserver;

public class MainActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        String listenPort = "8080";

        // Start the proxy server
        try {
            Proxyserver.StartProxy(listenPort);
        } catch (Exception e) {
            Log.e("ProxyServer", "Failed to start proxy: " + e.getMessage());
        }

        // Stop the proxy server (e.g., on a button click)
        findViewById(R.id.stopButton).setOnClickListener(v -> {
            try {
               Proxyserver.StopProxy();
            } catch (Exception e) {
                Log.e("ProxyServer", "Failed to stop proxy: " + e.getMessage());
            }
        });
    }
}
```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with clear and concise messages.
4. Submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact

For questions or support, please contact:

- **Telegram**: https://t.me/khaledagn
- **Instagram**: https://www.instagram.com/khaledagn
- **Youtube**: https://www.youtube.com/KhaledAGN
- **GitHub Issues**: [Create an Issue](https://github.com/khaledagn/go-http-proxy-server/issues)

---

Enjoy using the Go HTTP Proxy Server with VPN Sharing!