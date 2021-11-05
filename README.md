# log-output

Exercise 1.08

[https://github.com/pasiol/log-output/tree/1.08]

    pasiol@lab:~$ kubectl delete -f https://raw.githubusercontent.com/pasiol/log-output/1.07/manifests/ingress.yaml

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.08/manifests/service.yaml
    service/log-output-svc configured
    pasiol@lab:~$ kubectl get pods
    NAME                         READY   STATUS    RESTARTS   AGE
    log-output-6897c6f44-m22b2   1/1     Running   0          24m
    pasiol@lab:~$ kubectl logs log-output-6897c6f44-m22b2
    2021/11/05 14:40:12  Server started in port 8888.
    2021/11/05 14:41:11 request GET, /
    2021/11/05 14:41:11 67 bytes written
    2021-11-05T14:41:11.902300797Z d8599bed-11c8-40e1-adf0-d1659b6d556e
    2021/11/05 14:52:46 request GET, /
    2021/11/05 14:52:46 67 bytes written
    2021-11-05T14:52:46.102717301Z 908658fb-c4a3-4a90-a5b5-b494a526f72f
    2021/11/05 15:06:36 request GET, /
    2021/11/05 15:06:36 67 bytes written
    2021-11-05T15:06:36.082109454Z 6574988b-29f3-48e1-99eb-63c1ea632f94
    2021-11-05T15:06:36.35400004Z af354611-46b6-4851-a54c-06ff97c927eb
    2021/11/05 15:06:36 request GET, /favicon.ico
    2021/11/05 15:06:36 66 bytes written
    2021/11/05 15:06:39 request GET, /
    2021/11/05 15:06:39 67 bytes written
    2021-11-05T15:06:39.293338491Z 1529187f-83af-47d4-99ce-202a0b2a78be
