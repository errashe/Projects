#include <chrono>
#include <thread>
#include <future>

using namespace std;

float sync_func() {
	float res = 0;
	float mark = -1;

	for(float i=3; i<=10e6; i+=2) {
		res += mark * 4.0/i;
		mark = -mark;
	}

	return 4.0 + res;
}

float async_func(float start_point, float mark) {
	float res = 0;

	for(float i=start_point; i<=10e6; i+=4) {
		res += mark * 4.0/i;
	}

	return res;
}

int main() {
	auto t1 = chrono::high_resolution_clock::now();

	float res = 0;

	// for(int i = 0; i < 100; i++) {
		res = sync_func();

		// auto ret1 = async(&async_func, 3, -1);
		// auto ret2 = async(&async_func, 5, 1);

		// float i1 = ret1.get();
		// float i2 = ret2.get();

		// res = 4.0+i1+i2;
	// }

	printf("Result is %.30f\n", res);

	auto t2 = chrono::high_resolution_clock::now();

	std::chrono::duration<double, std::milli> fp_ms = (t2 - t1);

	printf("f() took %f ms\n", fp_ms.count());

	return 0;
}