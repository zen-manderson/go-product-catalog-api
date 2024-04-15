# Deployment Manifest for Charts

When it comes time to deploy your service. Your manifest will need to be configured for a gRPC service, rather than a HTTP. 
You can use this as a drop-in addition to your charts app. Be sure to change the name of `go-service-template` to your service.

For other assitence, you can utilize the following resources:
1. [Shipping a new service](https://counsl.atlassian.net/wiki/spaces/ENG/pages/3284140946/Building+a+new+service+or+application?atlOrigin=eyJpIjoiNzEyNmM4ZDU3MjY3NGM3MmFjMmZlZDNmNTczNzkyZTkiLCJwIjoiY29uZmx1ZW5jZS1jaGF0cy1pbnQifQ)
2. #engineering-devops Slack Channel
3. #guild-golang Slack Channel


## Example Application Manifest
```yaml
--- # don't forget to start a new array
_sharedConfig:
  appName: &name go-service-template
  containerPort: &containerPort 50050 # this is our standard grpc port
  opsPort: &opsPort 3001 # this is our standard health/metrics port
  cloudsql:
    enabled: false
  environmentVariables: &environment
    SERVICE_NAME: &name
    ENVIRONMENT: development
    HEALTH_ADDRESS: 0.0.0.0
    PORT: *containerPort
    OPS_PORT: *opsPort
    ENABLE_DEBUG_LOGGING: false
  image: &image
    pullPolicy: Always
    repository: us-central1-docker.pkg.dev/zen-dev-166315/zen-docker-images/go-service-template
    tag: d0129b8b1104ba793f4e45727ecf0ce29a2d7fa1
  labels: &labels
    zenbusiness.team/git-branch: main
    zenbusiness.team/git-repo: *name
    zenbusiness.team/pull-request-id: ""
  resources: &resources
    limits:
      memory: 128Mi
    requests:
      cpu: 70m
      memory: 128Mi
  serviceAccountName: &SAName go-service-tempalte@zen-dev-166315.iam.gserviceaccount.com # this is your services IAM Service Account
  volumeMounts: &extraMounts []
  volumes: &volumes []
autoscaling:
  horizontalPodAutoscalers:
    - name: *name
      targetDeployment: *name
      minReplicas: 1
      maxReplicas: 2
      targetCPUUtilizationPercentage: 80
  multidimPodAutoscalers: []
  verticalPodAutoscalers:
    - name: *name
      targetDeployment: *name
      updateMode: "Off"
configurations:
  configMaps: []
  externalSecrets:
    - name: *name
      secretStoreProviderId: gcpsm
  secretStores:
    - name: *name
      provider:
        gcpsm:
          projectID: zen-dev-secrets-2582
  secrets: []
networking:
  destinationRules:
    - host: *name
      name: *name
      subsets:
        - labels:
            zenbusiness.team/pull-request-id: ""
          name: *name
  services:
    - name: *name
      type: ClusterIP
      ports:
        - name: grpc-payments-api-service # service must be declared as a gRPC service
          containerPort: *containerPort
          servicePort: 443 # INFRA requires gRPC on port 443. Not port 80
          appProtocol: grpc
  virtualServices: [] # empty unless your service needs to be on the public gateway:
#    virtualServices:
#      - gateways:
#          - asm-ingress/asm-ingress-gateway
#        hosts:
#          - api.dev.zenbusiness.com
#        http:
#          - match:
#              - uri:
#                  prefix: /go-service-template/ # declare your services exposed prefix
#            name: *name
#            route:
#              - destination:
#                  host: *name
#                  port:
#                    number: 80
#                  subset: *name
#        name: *name
observability:
  podmonitoring:
    - name: go-service-template-metrics
      target: *name
      endpoints:
        - containerPort: *opsPort
security:
  authorizationPolicies:
    - name: go-service-template-allow-policy
      spec:
        action: ALLOW
        rules:
          - from:
              - source:
                  namespaces: ["go-service-template", "zenapi"] # which internal applications can talk to your service?
            to:
              - operation:
                  ports: ["50050"] # should match exposed container ports
          - to:
              - operation:
                  ports: ["3001"]
  peerAuthentication:
    - name: go-service-template-strict-mtls
      spec:
        mtls:
          mode: PERMISSIVE
        portLevelMtls:
          3001:
            mode: PERMISSIVE
        selector:
          matchLabels: *labels
  roleBindings:
    - name: pod-readers
      role: pod-reader
      subjects:
        - kind: ServiceAccount
          name: *name
  roles:
    - name: pod-reader
      rules:
        - apiGroups:
            - ""
          resources:
            - pods
          verbs:
            - get
            - watch
            - list
  serviceAccounts:
    - annotations:
        iam.gke.io/gcp-service-account: *SAName
      name: *name
storage:
  persistentVolumeClaims: []
workloads:
  cronJobs: []
  deployments:
    - annotations: {}
      autoscaling:
        enabled: true
      cloudsql:
        enabled: false
      environment:
      extraMounts: *extraMounts
      healthCheck:
        enabled: true
        httpGet:
          path: /alivez
          port: *opsPort
      image: *image
      labels: *labels
      name: *name
      podAntiAffinity:
        enabled: false
      topologySpreadConstraints:
        enabled: true
      ports:
        - containerPort: *containerPort
          name: internal-port
          protocol: TCP
      resources: *resources
      serviceAccountName: *name
      volumes: *volumes
  jobs: []
availability:
  podDisruptionBudgets:
    - name: *name
      maxUnavailable: 1
      labels: *labels

```
