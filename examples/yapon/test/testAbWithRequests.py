import json
import requests
import logging

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)


def request_api(url, params):

    try:
        r = requests.get(url, params=params, timeout=10)
        r.raise_for_status()  # 如果响应状态码不是 200，就主动抛出异常
    except requests.RequestException as e:
        print(e)
    else:
        json.dumps(r.json())
        return r.json(), r.elapsed.total_seconds()


def benchmark(num, url, params):

    for i in range(num):

        # total_time = 0
        json_text, req_time = request_api(url, params)

        # total_time += req_time
        # logger.info(
        #     '{} request, elapsed time {} {}'.format(
        #         i + 1, total_time, json_text))
        logger.info(
            '{} request'.format(i + 1))


if __name__ == '__main__':

    num = 500  # 压力测试次数
    url = 'http://192.168.88.158:8999/api/v1/ponds/stats/H19148841L'
    params = {
        "coilId": "H19148841L",
        "curDir": "E:/2250hrm/201911/20191108",
        "factorName": "wedge40",

        "functionName": "aimRate",
        "aim": 0,
        "tolerance": 20,
        "unit": "um",

        "lengthName": "s",
        "headLen": 17,
        "tailLen": 27,
        "headCut": 5,
        "tailCut": 5,

    }
    benchmark(num, url, params)
