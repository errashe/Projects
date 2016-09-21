org 100h

mov cx, 7
A:
push cx

mov cl, 4
mov bx, 1000
call beep

mov cl, 17
call delay

pop cx
dec cx
jnz A

mov cl, 120
mov bx, 1000
call beep

ret


;mov cl, 4
;mov bx, 1000
;call beep

;stop: jmp stop

; CL - long (x * 55ms)
delay:
  MOV DX, 40h
  MOV DS, DX
  MOV DL, [6Ch]
  ADD DL, CL

  .delay:
  CMP DL, [6Ch]
  JNZ .delay
ret

; BX - freq (1000)
; CL - long (40h)
beep:
  MOV AL, 10110110b
  OUT 43h, AL

  MOV AX, BX
  OUT 42h, AL
  MOV AL, AH
  OUT 42h, AL

  ; SOUND ON
  IN AL, 61h
  OR AL, 00000011b
  OUT 61h, AL

  ; WAIT
  call delay

  ; SOUND OFF
  IN AL, 61h
  AND AL, 11111100b
  OUT 61h, AL
ret










