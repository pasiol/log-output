# log-output

Exercise 1.07

[https://github.com/pasiol/log-output/tree/1.07]

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.07/manifests/deployment.yaml
    deployment.apps/log-output created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.07/manifests/service.yaml
    service/log-output-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.07/manifests/ingress.yaml
    ingress.networking.k8s.io/log-output-ingress created
    pasiol@lab:~$ kubectl get ing
    NAME                 CLASS    HOSTS   ADDRESS                            PORTS   AGE
    log-output-ingress   <none>   *       172.19.0.2,172.19.0.3,172.19.0.4   80      10s
    pasiol@lab:~$ kubectl get pods
    NAME                         READY   STATUS    RESTARTS   AGE
    log-output-6897c6f44-m22b2   1/1     Running   0          40s
    pasiol@lab:~$ kubectl logs log-output-6897c6f44-m22b2
    2021/11/05 14:40:12  Server started in port 8888.
    2021/11/05 14:41:11 request GET, /
    2021/11/05 14:41:11 67 bytes written
    2021-11-05T14:41:11.902300797Z d8599bed-11c8-40e1-adf0-d1659b6d556e
