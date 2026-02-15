FROM python:3.12-slim

WORKDIR /app

# 复制生成的 python 文件和 client.py
COPY python_gen/ ./python_gen

# 安装依赖
RUN pip install --no-cache-dir grpcio grpcio-tools

# 默认执行客户端
CMD ["python", "python_gen/client.py"]
