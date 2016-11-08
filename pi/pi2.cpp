#include <chrono>

#include <ctime>
#include <stdio.h>

#include <random>
#include <limits.h>

#include <omp.h>

#define N 5e6
#define THREADS 4

using namespace std;
using namespace std::chrono;

random_device rd;

float mrand() {
	return (float)(rd()) / (float)(UINT_MAX);
}

int nps(int count) {
	int np = 0;

	for(int i=0; i<count; i++) {
		float x = (float)mrand()*2-1;
		float y = (float)mrand()*2-1;

		if(x*x+y*y<=1) np++;
	}

	return np;
}

int main(int argc, char* argv[]) {
	srand(duration_cast<nanoseconds>(high_resolution_clock::now().time_since_epoch()).count());

	// future<int> handles[THREADS];

	// auto t1 = chrono::high_resolution_clock::now();
	// for(int i=0; i<THREADS; i++) {
	// 	handles[i] = async(&nps, N/THREADS);
	// }

	// int res = 0;
	// for(int i=0; i<THREADS; i++) {
	// 	res += handles[i].get();
	// }
	// auto t2 = chrono::high_resolution_clock::now();

	// printf("%f\n", 4.0f*res/N);


	// std::chrono::duration<double, std::milli> fp_ms = (t2 - t1);

	// printf("f() took %f ms\n", fp_ms.count());

	auto t1 = chrono::high_resolution_clock::now();

	int *iterationsOnThread = new int[THREADS];

	#pragma omp parallel num_threads(THREADS)
	{
		#pragma omp for
		for (int i = 0;i < THREADS;i++)
			iterationsOnThread[i] = nps(N/THREADS);
	}


	int res = 0;
	for(int i=0; i<THREADS; i++) {
		res += iterationsOnThread[i];
	}

	auto t2 = chrono::high_resolution_clock::now();

	std::chrono::duration<double, std::milli> fp_ms = (t2 - t1);

	printf("f() took %f ms\n", fp_ms.count());

	printf("%.5f\n", N);

	printf("%d\n", res);
	printf("%f\n", 4.0f*res/N);


	return 0;
}