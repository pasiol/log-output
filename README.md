# log-output

Exercise 1.01

    docker build -t pasiol/log-output .
    docker push pasiol/log-output

    pasiol@lab:~$ kubectl create deployment log-output --image=pasiol/log-output sha256:173888d2d767bd8992d62364c02b186e849984ff8f09ed4c91d49d8a3a631294
    deployment.apps/log-output created
    pasiol@lab:~$ kubectl get deployments
    NAME         READY   UP-TO-DATE   AVAILABLE   AGE
    log-output   1/1     1            1           14s
    pasiol@lab:~$ kubectl get pods
    NAME                          READY   STATUS    RESTARTS   AGE
    log-output-5fb6c5b797-r9tdk   1/1     Running   0          30s
    pasiol@lab:~$ kubectl logs log-output-5fb6c5b797-r9tdk
    2021-11-02T19:34:37.375961981Z f8042396-00bc-4204-a79a-6a5c96e95121
    2021-11-02T19:34:42.376966752Z d1b885e3-d0e0-4743-8978-8cd4820f9098
    2021-11-02T19:34:47.38112316Z bb33fab8-a4c1-48aa-abe0-d3d011bb4bac
    2021-11-02T19:34:52.385434229Z 19093f2f-0cbb-4946-8462-1bf14d940e24
    2021-11-02T19:34:57.389688899Z b568dd3c-b6a2-481d-b964-388b773674a4
    2021-11-02T19:35:02.393993297Z 3b0e04db-e0a9-4a02-8b3b-1725774d619d
    2021-11-02T19:35:07.394813376Z b8be12eb-2a13-4108-8c28-c689d58b0e46
    2021-11-02T19:35:12.399174911Z 1a768372-e963-4792-9056-95d361052076
    2021-11-02T19:35:17.400396817Z 22a5734f-a22e-4d42-81e1-fbccd7141b23
    2021-11-02T19:35:22.404653021Z 26c92cf6-c8f2-4518-9439-4fe4c11ef2f1
    2021-11-02T19:35:27.405144807Z 815b33c4-4bee-4950-911e-b3b8ad8113a1
