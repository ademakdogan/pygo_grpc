
import grpc
import concurrent
from concurrent import futures
import math_pb2
import math_pb2_grpc
from math_functions import sum_the_values, check_equality


class MathServicer(math_pb2_grpc.MathServicer):

    def OrderMath(self, request, context):
        print("Python script is running..")
        response = math_pb2.MathResponse()
        response.result = sum_the_values(request.x,request.y)
        response.message = check_equality(request.x,request.y)

        return response


def main():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers = 8))
    math_pb2_grpc.add_MathServicer_to_server(MathServicer(), server)
    print("Listening...")
    server.add_insecure_port('[::]:9001')
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    main()
