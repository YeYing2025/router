import React from 'react';
import { Tag } from 'antd';

const COLOR_MAP = {
  grey: 'default',
  black: 'default',
  olive: 'processing',
  violet: 'purple',
};

function AppTag({ className = '', children, basic, ...props }) {
  const nextClassName = ['router-ui-tag', className].filter(Boolean).join(' ');
  const nextColor = COLOR_MAP[props.color] || props.color;
  return (
    <Tag {...props} color={nextColor} className={nextClassName}>
      {children}
    </Tag>
  );
}

export default AppTag;
