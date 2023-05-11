import tkinter
import random
import time
import copy
import os.path
COORDINATEX=300
COORDINATEY=300
TURNCOUNTER=0
DIRECTION=0
SIDE=0
###
SPEED=0
TURNWILL=0
TIMER=0

ANCESTOR=open('A.mon.exe','rb')
DNA=ANCESTOR.read()
ANCESTOR.close()
def SPRITEGENERATOR():
    
    global SPEED
    SPEED=(((DNA[len(DNA)-1]+1)*30)//255)+1
    print('SPEED:',SPEED)

    global TURNWILL
    print(DNA[len(DNA)-2])
    TURNWILL=(((DNA[len(DNA)-2]+1)*100)//255)+1
    print('TURNWILL:',TURNWILL)

    global EYES
    EYES=[[[0,0,0,0],0,0],[[0,0,0,0],0,0]]
    EYES[0][0][0]=15
    EYES[0][0][1]=15
    EYES[0][0][2]=18
    EYES[0][0][3]=18
    EYES[0][1]='#{:02x}{:02x}{:02x}'.format((((DNA[len(DNA)-3]+1)*255)//255)+1,(((DNA[len(DNA)-4]+1)*255)//255)+1,(((DNA[len(DNA)-5]+1)*255)//255)+1  )
    EYES[0][2]='#{:02x}{:02x}{:02x}'.format((((DNA[len(DNA)-6]+1)*255)//255)+1,(((DNA[len(DNA)-7]+1)*255)//255)+1,(((DNA[len(DNA)-8]+1)*255)//255)+1  )
    EYES[1][0][0]=25
    EYES[1][0][1]=15
    EYES[1][0][2]=28
    EYES[1][0][3]=18
    EYES[1][1]='#{:02x}{:02x}{:02x}'.format((((DNA[len(DNA)-9]+1)*255)//255)+1,(((DNA[len(DNA)-10]+1)*255)//255)+1,(((DNA[len(DNA)-5]+11)*255)//255)+1  )
    EYES[1][2]='#{:02x}{:02x}{:02x}'.format((((DNA[len(DNA)-12]+1)*255)//255)+1,(((DNA[len(DNA)-13]+1)*255)//255)+1,(((DNA[len(DNA)-14]+1)*255)//255)+1  )

    global BODY
    BODY=[]

    #for X in range((((DNA[len(DNA)-15]+1)*10)//255)+1):
    for X in range(5):
        BODY.append([])

    TRACKDNA=16
    for X in range(len(BODY)):
        BODY[X].append([])
        TRACKDNA+=1

    for X in range(len(BODY)):
        BODY[X].append('#{:02x}{:02x}{:02x}'.format(  (((DNA[len(DNA)-(TRACKDNA+X)]+1)*255)//255)+1,(((DNA[len(DNA)-(TRACKDNA+X+1)]+1)*255)//255)+1,(((DNA[len(DNA)-(TRACKDNA+X+2)]+1)*255)//255)+1,  ))
        TRACKDNA+=3                   

    TRACKDNA+=1
    OTHERDNATRACKER=TRACKDNA

    for X in range(len(BODY)):
        for Y in range( ((((DNA[len(DNA)-(TRACKDNA)]+1)*7)//255)+1)*2 ):        
            BODY[X][0].append(     (((DNA[len(DNA)-(OTHERDNATRACKER+X)]+1)*32)//255)+1         )
            BODY[X][0].append(     (((DNA[len(DNA)-(OTHERDNATRACKER+X+1)]+1)*32)//255)+1          )
            OTHERDNATRACKER+=2              
    print('BODY:',BODY)    
    

SPRITEGENERATOR()
    
       

def SPRITESHOW():
    def SPRITEACTION():
        #LOADS GLOBAL VARIABLES
        global COORDINATEX
        global COORDINATEY
        global TURNCOUNTER
        global DIRECTION
        global SIDE
        global BODY
        global TIMER

        #WIGGLES POLYGON
        IDLEBODY=copy.deepcopy(BODY)
        for X in IDLEBODY:
            for V in range(len(X[0])):
                X[0][V]+=random.randint(-1,1)
        
        #INCREASES TURNCOUNTER
        if TURNCOUNTER>=1000:
            DIRECTION=random.randint(0,7)
            TURNCOUNTER=0
        TURNCOUNTER+=random.randint(TURNWILL-10,TURNWILL+10)

        
        #CHANGES DIRECTION AND SIDE
        if DIRECTION==0 and COORDINATEX>0:
            COORDINATEX-=1
            SIDE=0
        if DIRECTION==1 and COORDINATEX<SPRITE.winfo_screenwidth():
            COORDINATEX+=1
            SIDE=1
        if DIRECTION==2 and COORDINATEY<SPRITE.winfo_screenheight():
            COORDINATEY+=1
        if DIRECTION==3 and COORDINATEY>0:
            COORDINATEY-=1
        if DIRECTION==4 and COORDINATEX>0 and COORDINATEY>0:
            COORDINATEX-=1
            COORDINATEY-=1
            SIDE=0
        if DIRECTION==5 and COORDINATEX<SPRITE.winfo_screenwidth() and COORDINATEY>0:
            COORDINATEX+=1
            COORDINATEY-=1
            SIDE=1
        if DIRECTION==6 and COORDINATEY<SPRITE.winfo_screenheight() and COORDINATEX<SPRITE.winfo_screenwidth():
            COORDINATEY+=1
            COORDINATEX+=1
            SIDE=1
        if DIRECTION==7 and COORDINATEX>0 and COORDINATEY<SPRITE.winfo_screenheight():
            COORDINATEY+=1                   
            COORDINATEX-=1
            SIDE=0

        #DELETES PREVIOUS FRAME AND CREATES NEW WINDOW
        SCREEN.delete('all')        
        SPRITE.geometry('{}x{}+{}+{}'.format(32,32, COORDINATEX,COORDINATEY))
      
        #BODY
        if SIDE==0:
            for I in IDLEBODY:
                SCREEN.create_polygon(I[0],fill=I[1])
        if SIDE==1:            
            for I in IDLEBODY:
                MIRROR=[]
                COUNTER=0
                for M in I[0]:
                    if COUNTER%2 == 0:     
                        MIRROR.append(32-M)
                    else:
                        MIRROR.append(M)
                    COUNTER+=1 
                SCREEN.create_polygon(MIRROR,fill=I[1])

        #EYES
        if SIDE==0:
            for E in EYES:
                SCREEN.create_oval(E[0],fill=E[1],outline=E[2],width=1)
        if SIDE==1:
            for E in EYES:
                MIRROR=[]
                COUNTER=0
                for M in E[0]:
                    if COUNTER%2 == 0:                        
                        MIRROR.append(32-M)
                    else:
                        MIRROR.append(M)
                    COUNTER+=1                    
                SCREEN.create_oval(MIRROR,fill=E[1],outline=E[2],width=1)

        #PLAY FRAMES
        SPRITE.after(random.randint((SPEED-3)+(SPEED-3)**4//8000,(SPEED+3)+(SPEED+3)**4//8000),SPRITEACTION)
       
    SPRITE=tkinter.Tk()
    SPRITE.overrideredirect(True)
    SPRITE.wm_attributes('-transparentcolor','pink')
    SCREEN=tkinter.Canvas(SPRITE, bg='pink', highlightthickness=0)
    SCREEN.pack()
    SCREEN.after(0,SPRITEACTION)
    SPRITE.after(60000, SPRITE.destroy)
    SPRITE.mainloop()

SPRITESHOW()
