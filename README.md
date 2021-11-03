# log-output

Exercise 1.03

[https://github.com/pasiol/log-output/tree/1.03]

    pasiol@lab:~$ kubectl apply -f https://raw.githubusercontent.com/pasiol/log-output/1.03/manifests/deployment.yaml
    deployment.apps/log-output created
    pasiol@lab:~$ kubectl get deployments
    NAME         READY   UP-TO-DATE   AVAILABLE   AGE
    log-output   1/1     1            1           6s
    pasiol@lab:~$ kubectl get pods
    NAME                          READY   STATUS    RESTARTS   AGE
    log-output-6cb768654c-tnnw5   1/1     Running   0          14s
    pasiol@lab:~$ kubectl logs log-output-6cb768654c-tnnw5
    2021-11-03T15:48:31.16089871Z feb100cc-2e39-4533-99a9-04c88606d1b8
    2021-11-03T15:48:36.16514Z ac343023-1053-4f1c-8376-ceb1b25982d6
    2021-11-03T15:48:41.169377861Z 03b1293d-beab-470e-9e16-5456c4c02cca
    2021-11-03T15:48:46.169808997Z 37f2943a-ce7e-498a-b0f5-3193f43ffaa9
    2021-11-03T15:48:52.654197224Z 843ebacd-2b60-48ad-950d-5e3351366f1f
    2021-11-03T15:48:57.654731762Z cd81cf97-6597-431b-847e-8be7958f16fb
    2021-11-03T15:49:02.65895619Z 0b4e360f-a4fd-4160-8fb1-39fa7115511a
    2021-11-03T15:49:07.663353747Z 544a7cf2-948e-4f36-a775-cacba09e07d0
    2021-11-03T15:49:12.664039798Z da184384-3072-477f-a022-5bb02c83adc0
