# Go Learning Project

This repository is designed for **beginner-friendly, hands-on learning of Go (Golang)**.  
It includes practical examples covering core Go features, concurrency, HTTP servers, testing, and using external libraries.

---

## **Installation (Linux)**

### **Step 1: Install Go**

```bash
# Download latest Go tarball
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz

# Remove any previous installation
sudo rm -rf /usr/local/go

# Extract Go
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz

# Add Go to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile

# Verify installation
go version
```
