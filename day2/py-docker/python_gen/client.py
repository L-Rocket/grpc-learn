import grpc
import user_pb2, user_pb2_grpc

def main():
    channel = grpc.insecure_channel("grpc-server:50051")
    stub = user_pb2_grpc.UserServiceStub(channel)

    req = user_pb2.GetUserRequest(id=123)
    resp = stub.GetUser(req)

    print("Python client 收到响应:")
    print("name =", resp.name)
    print("age  =", resp.age)

if __name__ == "__main__":
    main()
