# Welcome to Tilt!
#   To get you started as quickly as possible, we have created a
#   starter Tiltfile for you.
#
#   Uncomment, modify, and delete any commands as needed for your
#   project's configuration.


# Output diagnostic messages
#   You can print log messages, warnings, and fatal errors, which will
#   appear in the (Tiltfile) resource in the web UI. Tiltfiles support
#   multiline strings and common string operations such as formatting.
#
#   More info: https://docs.tilt.dev/api.html#api.warn
print("""
-----------------------------------------------------------------
✨ Hello Tilt! This appears in the (Tiltfile) pane whenever Tilt
   evaluates this file.
-----------------------------------------------------------------
""".strip())
warn('ℹ️ Open {tiltfile_path} in your favorite editor to get started.'.format(
    tiltfile_path=config.main_path))

# Create namespace
local_resource(
    'create-namespace',
    cmd='kubectl create namespace astarte-kubernetes-operator-system || true',
    deps=[],
)

# Install Cert-Manager
local_resource(
    'install-cert-manager',
    cmd='kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.0/cert-manager.yaml',
    deps=[],
    auto_init=True,
    trigger_mode=TRIGGER_MODE_MANUAL
)

# Delete existing CRDs
local_resource(
    'delete-crds',
    cmd='''
    kubectl delete crd astartedefaultingresses.ingress.astarte-platform.org --ignore-not-found
    kubectl delete crd astartes.api.astarte-platform.org --ignore-not-found
    kubectl delete crd flows.api.astarte-platform.org --ignore-not-found
    ''',
)

# Build and push the operator image
docker_build(
    'controller',
    '.',
    dockerfile='./Dockerfile',
    only=[
        './controllers',
        './apis',
        './lib',
        'go.mod',
        'go.sum'
    ]
)

# Create service account
local_resource(
    'create-sa',
    cmd='kubectl create serviceaccount astarte-kubernetes-operator-controller-manager -n astarte-kubernetes-operator-system || true',
    deps=[],
)

# Create certificates directory, generate TLS certs and create secret
local_resource(
    'setup-certs',
    cmd='''
    mkdir -p /tmp/k8s-webhook-server/serving-certs && 
    openssl req -x509 -newkey rsa:4096 -keyout /tmp/k8s-webhook-server/serving-certs/tls.key -out /tmp/k8s-webhook-server/serving-certs/tls.crt -days 365 -nodes -subj "/CN=webhook-service.default.svc" &&
    kubectl create secret tls webhook-server-cert --cert=/tmp/k8s-webhook-server/serving-certs/tls.crt --key=/tmp/k8s-webhook-server/serving-certs/tls.key -n astarte-kubernetes-operator-system || true
    ''',
    deps=[],
)

# Install fresh CRDs
local_resource(
    'install-crds',
    cmd='make install',
    deps=['config/crd/bases'],
    resource_deps=['delete-crds'],
    auto_init=True,
    trigger_mode=TRIGGER_MODE_MANUAL
)

# Deploy using kustomize
k8s_yaml(kustomize('config/default'))

# Configure the operator resource
k8s_resource(
    'astarte-kubernetes-operator-controller-manager',
    port_forwards=['8080:8080'],
    auto_init=True,
    trigger_mode=TRIGGER_MODE_AUTO,
    resource_deps=['create-namespace', 'setup-certs', 'install-crds']
)

# Watch for changes
watch_file('config')
watch_file('controllers')
watch_file('apis')
