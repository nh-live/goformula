DRAWTEXT(CLOSE/OPEN>1.08,LOW,'abc'), COLORBLUE;
DRAWLINE(HIGH>=HHV(HIGH,20),HIGH,LOW<=LLV(LOW,20),LOW,1), COLORRED;
PLOYLINE(HIGH>=HHV(HIGH,20),HIGH),COLORRED,LINETHICK9;
STICKLINE(CLOSE>OPEN,CLOSE,OPEN,0.8,1), COLORCYAN;
DRAWICON(CLOSE>OPEN,LOW,1);
DRAWKLINE(HIGH,OPEN,LOW,CLOSE), NODRAW;