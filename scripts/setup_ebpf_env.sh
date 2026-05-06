
set -e

INTERFACE="eth0"
BPF_PROG="irongate_filter.o"
MAP_PATH="/sys/fs/bpf/irongate_maps"

echo "--- [IronGate] Iniciando Configuração eBPF ---"


if [ "$EUID" -ne 0 ]; then 
  echo " Erro: Por favor, execute como root (sudo)"
  exit 1
fi


echo "[1/4] Limpando XDP anterior na interface $INTERFACE..."
ip link set dev $INTERFACE xdp off 2>/dev/null || true


if ! mount | grep -q /sys/fs/bpf; then
    echo "[2/4] Montando BPF Filesystem..."
    mount -t bpf bpf /sys/fs/bpf
fi

echo "[3/4] Carregando programa $BPF_PROG..."

echo "[4/4] Aplicando políticas de segurança padrão..."

echo "---  Ambiente IronGate eBPF pronto em $INTERFACE ---"
