BYTES = open('NAME OF FILE TO BE EXTRACTED', "rb")
data = BYTES.read()
print(len(data))
EXTRACTION=[]
for X in data:
    EXTRACTION.append(str(int(X)))
print(EXTRACTION)
