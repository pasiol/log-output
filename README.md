# log-output

Exercise 1.10

[https://github.com/pasiol/log-output/tree/1.10]

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.10/manifests/deployment.yaml
    deployment.apps/log-output-dep created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.10/manifests/service.yaml
    service/log-output-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.10/manifests/ingress.yaml
    ingress.networking.k8s.io/log-output-ingress created
    pasiol@lab:~$ kubectl get pods
    NAME                              READY   STATUS    RESTARTS   AGE
    ping-pong-5958c444d8-zqjcm        1/1     Running   3          19h
    log-output-dep-549d788b8f-blbzx   2/2     Running   0          24s
    pasiol@lab:~$ kubectl logs log-output-dep-549d788b8f-blbzx log-output-writer
    2021/11/07 11:46:19 starting writer
    pasiol@lab:~$ kubectl logs log-output-dep-549d788b8f-blbzx log-output-reader
    2021/11/07 11:46:19 starting reader
    2021/11/07 11:46:19 starting in address 0.0.0.0:3000.


