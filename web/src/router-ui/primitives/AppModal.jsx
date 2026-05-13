import React from 'react';
import { Modal } from 'antd';

function mapModalWidth(size) {
  switch (size) {
    case 'tiny':
      return 420;
    case 'small':
      return 560;
    case 'large':
      return 960;
    case 'fullscreen':
      return '100vw';
    default:
      return 720;
  }
}

function AppModal({
  className = '',
  children,
  open,
  onClose,
  title,
  footer,
  size,
  closeOnDimmerClick = true,
  ...props
}) {
  const nextClassName = ['router-ui-modal', className].filter(Boolean).join(' ');
  return (
    <Modal
      {...props}
      className={nextClassName}
      open={open}
      onCancel={onClose}
      title={title}
      footer={footer ?? null}
      width={mapModalWidth(size)}
      mask={{ closable: closeOnDimmerClick }}
    >
      {children}
    </Modal>
  );
}

export default AppModal;
