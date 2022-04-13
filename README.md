

### kubecm



指令

- version

- help

- init

    ```
    mkdir ~/.kube/kubeconfig
    touch ~/.kube/kubecm.yaml
    ```

    

- add

    ```
    kubecm add foo -f config.xxx --move
    ```

​	

​	将 -f config.xxx 移动到 ~/.kube/kubeconfig

​	ln -sf ~/.kube/kubeconfig config

​	修改 ~/.kube/kubecm.yaml .current: foo

​	增加 foo: ~/.kube/kubeconfig/config.xxx



- del

    ```
    kubecm del foo
    ```

    

- desc

    ```
    kubem desc alias
    
    # contents of config 
    ```

    

- list

    ```
    > kubecm list 
    Alias									Location
    foo1				 					~/.kube/kubeconfig/config.foo1
    foo2 (current)				~/.kube/kubeconfig/config.foo2  # color: green, bold
    foo3				 					~/.kube/kubeconfig/config.foo3
    
    ```

    

- switch

    ```
    kubecm switch # 进入交互选择界面
    
    kubecm switch alias 直接切换
    ```

    

- rename

    ```
    kubecm rename foo1 eks-foo
    ```

    



> ~/.kubecm/config

```
path: ~/.kube/kubeconfig
current: tke-persona
alias: 
  eks-alopoker: config.alopoker-eks
  eks-bingo: config.bingo-eks
  tke-persona: config.tke-persona
```



> 

```
config -> kubeconfig/config.tke-persona
kubeconfig
├── config.alopoker-eks
├── config.bingo-eks
├── config.bingo-tke
├── config.cicd-local
├── config.cicd.eks-dev-pwmrnyjq
├── config.cooking-tke
├── config.dango-tke
├── config.eu-west-2.makerx-london-eks
├── config.local-dev
├── config.microk8s-10.192.0.2
├── config.planetparty-eks
├── config.summonquest-eks
├── config.tke-cooking
├── config.tke-persona
├── config.us-west-2.cheesecake-k8s
├── config.us-west-2.cookingbestie-eks
├── config.us-west-2.dreamlife-eks-cluster
├── config.us-west-2.fantasy-eks-cluster
├── config.us-west-2.farmtopia-aklucpcl
├── config.us-west-2.frozencity-k8s
├── config.us-west-2.k8s-devops
├── config.us-west-2.kachingcasino
├── config.us-west-2.makerX-global-eks
├── config.us-west-2.mergeworld-k8s
└── config.us-west-2.zhongtai-cluster
```

