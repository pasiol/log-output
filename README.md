# log-output

Exercise 1.10

[https://github.com/pasiol/log-output/tree/1.10]

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.10/manifests/deployment.yaml
    deployment.apps/log-output-dep created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.10/manifests/service.yaml
    service/log-output-svc created
    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.10/manifests/ingress.yaml
    ingress.networking.k8s.io/log-output-ingress created
    pasiol@lab:~$ kubectl get ingress
    NAME                 CLASS    HOSTS   ADDRESS                            PORTS   AGE
    log-output-ingress   <none>   *       172.19.0.2,172.19.0.3,172.19.0.4   80      7s
    pasiol@lab:~$ curl http://172.19.0.2
    2021-11-07T14:37:32.341154859Z cb616992-c96b-4c97-86cf-dc97713f9ac0
    2021-11-07T14:37:37.345494932Z cd0ad1b6-1575-40b6-bb5c-4838a74ffb38
    2021-11-07T14:37:42.348028001Z 972490ce-9d84-498f-9613-caad1e6b73d6
    2021-11-07T14:37:47.352435027Z 0a689f0b-5ba2-4ed7-9dd5-7facf175dae8
    2021-11-07T14:37:52.354157632Z cb1964c1-6fbd-4248-a2ba-f3e9502e0ea1
    2021-11-07T14:37:57.356200019Z bb6a2f80-9cd0-4fe4-8ed4-a19ccf698b37
    2021-11-07T14:38:02.360735682Z 5c6fc3e1-0c0c-42dd-823d-87c6457f4147
    2021-11-07T14:38:07.36520786Z 85723394-dcfd-4643-8ab1-e0592b80d4ea
    2021-11-07T14:38:12.369662629Z bb6efa16-183e-47de-8fcb-e6e50b64b63c
    2021-11-07T14:38:17.371134268Z e13ea37f-8d03-40fb-87e1-a033ef4fd28c
    pasiol@lab:~$ kubectl get pods
    NAME                              READY   STATUS    RESTARTS   AGE
    log-output-dep-6fd5bcfd55-zrjxf   2/2     Running   0          61s
    pasiol@lab:~$ kubectl logs log-output-dep-6fd5bcfd55-zrjxf log-output-writer
    2021/11/07 14:37:32 starting writer
    pasiol@lab:~$ kubectl logs log-output-dep-6fd5bcfd55-zrjxf log-output-reader
    2021/11/07 14:37:35 starting reader
    2021/11/07 14:37:35 starting in address 0.0.0.0:3000.
    2021/11/07 14:38:21 getting request from 10.42.1.4:41302
    2021/11/07 14:38:21 readed file /var/app/data/uuids.txt
    2021/11/07 14:38:21 679 bytes written 10.42.1.4:41302



