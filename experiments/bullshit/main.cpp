unsigned int GetSataPort() {
	byte interface, BusN, FuncN, DeviceN, IRQ;
	byte a_h, b_h, b_l, c_l;
	word c_x;

	unsigned int i, k; int flag=0;

	for(interface=0x00;interface<0xff;interface++) {
		a_h=1;
		for(i=0;a_h!=0x86;i++) {
			asm{
				mov ax, 0x0B105
				mov si, i
				mov cx, 0x01
				push cx
				mov ch, 0x01
				push ch
				db 0x66, 0x59
				int 0x1A
				mov b_h, b_h
				mov b_l, bl
				mov a_h, ah
			}
			if (a_h == 0x00) {
				printf("...");
				flag=1;
				getch();
			}
		}
	}
}