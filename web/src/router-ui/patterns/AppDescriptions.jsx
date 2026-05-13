import React from 'react';
import { Descriptions } from 'antd';

const normalizeItems = (items) =>
  (Array.isArray(items) ? items : []).map((item) => ({
    key: item?.key,
    label: item?.label,
    children: item?.children ?? item?.value ?? '-',
    span: item?.span,
  }));

function AppDescriptions({
  className = '',
  items = [],
  bordered = true,
  size = 'small',
  column = { xs: 1, sm: 1, md: 2, lg: 2, xl: 2 },
  ...props
}) {
  const nextClassName = ['router-ui-descriptions', className]
    .filter(Boolean)
    .join(' ');
  return (
    <Descriptions
      {...props}
      className={nextClassName}
      bordered={bordered}
      size={size}
      column={column}
      items={normalizeItems(items)}
    />
  );
}

export default AppDescriptions;
