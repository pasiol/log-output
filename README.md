# log-output

Exercise 1.08

[https://github.com/pasiol/log-output/tree/1.08]

    pasiol@lab:~$ kubectl delete -f https://raw.githubusercontent.com/pasiol/log-output/1.07/manifests/ingress.yaml

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.08/manifests/service.yaml
    service/log-output-svc configured
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.08/manifests/deployment.yaml
    deployment.apps/log-output created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.08/manifests/service.yaml
    service/log-output-svc created
    pasiol@lab:~$ kubectl get pods
    NAME                          READY   STATUS    RESTARTS   AGE
    log-output-5ff9857984-n7tbg   1/1     Running   0          72s
    pasiol@lab:~$ kubectl logs log-output-5ff9857984-n7tbg
    2021/11/06 13:29:24 server started in port 8888.
    2021-11-06T13:40:23.182029808Z 56145b97-f009-4ae8-82ba-7a7e7bf97976
    2021/11/06 13:40:23 request GET, /, 127.0.0.1:34338
    2021/11/06 13:40:23 67 bytes written 127.0.0.1:34338
