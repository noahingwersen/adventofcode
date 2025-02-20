def opcode(reg, prog, verbose=False, part=1):
    i = 0
    output = []

    while True:    
        op = prog[i]
        li = prog[i+1]
        
        # combo from literal
        co = li 
        if co>3 and co<7:
            co=reg[co-4]        
        
        if verbose:
            print(f"{i} | {op} : {co} | ",end="")
        
        if op==0: # adv
            reg[0] = reg[0] // 2**co  
        elif op==1: # bxl
            reg[1] = reg[1] ^ li 
        elif op==2: # bst
            reg[1] = co % 8
        elif op==3: # jnz
            if reg[0]!=0: 
                i = li - 2
        elif op==4: # bxc
            reg[1] = reg[1] ^ reg[2]
        elif op==5: # out
            output += [ co % 8 ]
        elif op==6: # bdv
            reg[1] = reg[0] // 2**co  
        elif op==7: # cdv
            reg[2] = reg[0] // 2**co  

        if verbose:
            print(reg)

        i+=2
        if i>=len(prog):
            break
    return output,reg

def findA(prog,a=0,b=0,c=0,ip=-1):
    if abs(ip) > len(prog): 
      return a
    for i in range(8):
        aa = a * 8 + i
        reg = [ aa , b, c ]
        output, reg = opcode(reg,prog)
        if output[0]==prog[ip]:
            aa = findA(prog, aa, reg[1], reg[2], ip-1) 
            # backtracking here in case next octal cannot be found (loop is not broken as in greedy solution)
            if aa:
                return aa
    return None



print(findA([2,4,1,7,7,5,0,3,1,7,4,1,5,5,3,0]))