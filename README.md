# google translate ip validity verification


example：

gt_check https://cdn.jsdelivr.net/npm/@hcfy/google-translate-ip/ips.txt

or

gt_check -f ipfile.txt

or

gt_check ip1 ip2 ip3 ....


![image](https://user-images.githubusercontent.com/11746715/224201611-7390b007-fa1a-4a7d-b4b2-96e68323f378.png)

- select some ip with a delay of less than 300ms

- edit hosts file (C:\Windows\System32\drivers\etc\hosts)

- add the selected ip address at the end of the file

```bash
142.251.12.90        translate.googleapis.com
142.251.8.90         translate.googleapis.com
142.250.157.184      translate.googleapis.com
108.177.97.100       translate.googleapis.com
142.250.157.186      translate.googleapis.com

142.251.12.90        translate.google.com
142.251.8.90         translate.google.com
142.250.157.184      translate.google.com
108.177.97.100       translate.google.com
142.250.157.186      translate.google.com
```

- save file!

- cmd ipconfig /flushdns

- open your browser and test it...
![image](https://user-images.githubusercontent.com/11746715/224202889-b9eca04a-7e8c-4098-8a53-7f88905592ce.png)



