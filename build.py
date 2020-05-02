import re,os,shutil,urllib.request
shutil.copytree("WebPage","WebPageBuild")
f=open("WebPageBuild/index.html","r",encoding='utf-8')
content=f.read()
link=re.findall('http(?:.+?)(?=\")',content)
f.close()
opener = urllib.request.build_opener()
opener.addheaders = [('User-agent', 'Opera/9.80 (Android 2.3.4; Linux; Opera Mobi/build-1107180945; U; en-GB) Presto/2.8.149 Version/11.10')]
urllib.request.install_opener(opener)
for i in link:
    a=i.split(r"/")[-1]
    print(a)
    urllib.request.urlretrieve(i,"WebPageBuild/"+a)
    content=content.replace(i,a)
with open("WebPageBuild/index.html","w",encoding='utf-8') as f2:
    f2.write(content)
f2.close()
os.system('go-bindata -fs -prefix "WebPageBuild/" WebPageBuild/')
os.system("go-bindata-assetfs WebPageBuild/...")
os.system("go build")
shutil.rmtree("WebPageBuild")
