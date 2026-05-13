import React from 'react';
import { Table } from 'antd';
import AppEmpty from './AppEmpty';

function AppTable({
  className = '',
  size = 'small',
  pagination = false,
  locale,
  ...props
}) {
  const nextClassName = ['router-ui-table', className].filter(Boolean).join(' ');
  const nextLocale = {
    emptyText: <AppEmpty>-</AppEmpty>,
    ...(locale || {}),
  };

  return (
    <Table
      {...props}
      className={nextClassName}
      size={size}
      pagination={pagination}
      locale={nextLocale}
    />
  );
}

export default AppTable;
