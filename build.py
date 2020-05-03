import re,os,shutil,urllib.request
shutil.copyfile("main.go","main.go.org")
shutil.copytree("WebPage","WebPageBuild")
f=open("WebPageBuild/index.html","r",encoding='utf-8')
content=f.read()
link=re.findall('http(?:.+?)(?=\")',content)
f.close()
for i in link:
    a=i.split(r"/")[-1]
    print(a)
    urllib.request.urlretrieve(i,"WebPageBuild/"+a)
    content=content.replace(i,a)
with open("WebPageBuild/index.html","w",encoding='utf-8') as f2:
    f2.write(content)
f2.close()
f3=open("main.go","r",encoding='utf-8')
t=f3.read()
f3.close()
t=t.replace('http.Dir("./WebPage/")','assetFS()')
with open("main.go","w",encoding='utf-8') as f4:
    f4.write(t)
f4.close()
os.system('go-bindata -fs -prefix "WebPageBuild/" WebPageBuild/')
os.system("go-bindata-assetfs WebPageBuild/...")
os.system("go build")
os.remove("main.go")
os.rename("main.go.org","main.go")
shutil.rmtree("WebPageBuild")
