# # Menggunakan base image Golang
# FROM golang:latest

# # Set working directory dalam container
# WORKDIR /app

# # Copy semua file dari local ke dalam container
# COPY . .

# # Jalankan perintah untuk mendownload dependencies (jika ada)
# RUN go mod tidy
# RUN cp .env.production .env
# RUN go build -o ./build/api/main.exe ./cmd/api/main.go

# # Ekspos port (opsional, hanya jika aplikasi menggunakan port tertentu)
# EXPOSE 8080

# Perintah default untuk menjalankan aplikasi saat container berjalan
# CMD ["go build -o ./build/api/main.exe ./cmd/api/main.go"]
# CMD ["./build/api/main.exe"]

# Gunakan base image
FROM golang:latest

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . . 

# Hapus file .env jika ada konflik
# RUN rm -f .env

# Install dependencies
RUN go mod tidy

# Ganti file .env dengan versi production
# RUN cp .env.production .env

# Build aplikasi
RUN go build -o ./build/api/main.exe ./cmd/api/main.go

# Ekspos port
EXPOSE 8080

# Perintah default untuk menjalankan aplikasi
CMD ["./build/api/main.exe"]

