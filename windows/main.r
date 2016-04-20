png("plot.png", width = 10, height = 40, units = 'in', res = 100)

draw_fft <- function(fft_dots, name) {
	plot(a, abs(fft_dots), "h"); title(name)
}

draw_graph <- function(dots, name) {
	plot(a, dots, "l"); title(name)
}

barlett <- function(N) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- 1 - abs(2*n/(N-1) - 1)
		r <- c(r, temp)
	}
	return(r)
}

hann <- function(N) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- 0.5 - 0.5*cos(2*pi*n/(N-1))
		r <- c(r, temp)
	}
	return(r)
}

hamming <- function(N) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- 0.54 - 0.46*cos(2*pi*n/(N-1))
		r <- c(r, temp)
	}
	return(r)
}

blackman <- function(N) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- 0.42 - 0.5*cos(2*pi*n/(N-1))+0.08*cos(4*pi*n/(N-1))
		r <- c(r, temp)
	}
	return(r)
}

natoll <- function(N) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- 0.3635819 - 0.4891775*cos(2*pi*n/(N-1)) + 0.1365995*cos(4*pi*n/(N-1)) - 0.0106411*cos(6*pi*n/(N-1))
		r <- c(r, temp)
	}
	return(r)
}

gauss <- function(N, a) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- exp(-1/2 * (2*a*n/(N-1) - a)^2)
		r <- c(r, temp)
	}
	return(r)
}

I_0 <- function(x) {
	r <- 0
	for(k in seq(1, 32, 1)) {
		r <- r + (x/2)^k/factorial(k)
	}
	return(r^2)
}

kayzer_bessel <- function(N, a) {
	r <- NULL
	for(n in seq(1, N, 1)) {
		temp <- I_0(pi*a*(1 - ((n-N/2)/(N-2))^2)^0.5)/I_0(pi*a)
		r <- c(r, temp)
	}
	return(r)
}

a <- seq(-10, 10, 0.1)
b <- sin(2*a) * sin(a/2)

n <- length(b)
cnt <- 0

f1 <- b; 														fft1 <- fft(f1); cnt <- cnt + 1
f2 <- b * blackman(n); 							fft2 <- fft(f2); cnt <- cnt + 1
f3 <- b * natoll(n);								fft3 <- fft(f3); cnt <- cnt + 1
f4 <- b * barlett(n);								fft4 <- fft(f4); cnt <- cnt + 1
f5 <- b * hann(n);									fft5 <- fft(f5); cnt <- cnt + 1
f6 <- b * hamming(n);								fft6 <- fft(f6); cnt <- cnt + 1
f7 <- b * gauss(n, 5);							fft7 <- fft(f7); cnt <- cnt + 1
f8 <- b * kayzer_bessel(n, 100);		fft8 <- fft(f8); cnt <- cnt + 1

par(mfrow = c(cnt,2))
draw_graph(f1, "Исходный график"); 									draw_fft(fft1, "Преобразование фурье")
draw_graph(f2, "Исходный график (blackman)");				draw_fft(fft2, "Преобразование фурье (blackman)")
draw_graph(f3, "Исходный график (natoll)");					draw_fft(fft3, "Преобразование фурье (natoll)")
draw_graph(f4, "Исходный график (barlett)");				draw_fft(fft4, "Преобразование фурье (barlett)")
draw_graph(f5, "Исходный график (hann)");						draw_fft(fft5, "Преобразование фурье (hann)")
draw_graph(f6, "Исходный график (hamming)");				draw_fft(fft6, "Преобразование фурье (hamming)")
draw_graph(f7, "Исходный график (gauss)");					draw_fft(fft7, "Преобразование фурье (gauss)")
draw_graph(f8, "Исходный график (kayzer_bessel)");	draw_fft(fft8, "Преобразование фурье (kayzer_bessel)")
