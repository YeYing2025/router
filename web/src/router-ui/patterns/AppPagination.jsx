import React from 'react';
import { Pagination } from 'antd';

function AppPagination({
  className = '',
  activePage,
  current,
  totalPages,
  onPageChange,
  onChange,
  ...props
}) {
  const resolvedCurrent = Number(current || activePage || 1) || 1;
  const resolvedTotalPages = Number(totalPages || 1) || 1;
  const nextClassName = ['router-ui-pagination', className]
    .filter(Boolean)
    .join(' ');

  const handleChange = (page, pageSize) => {
    if (typeof onChange === 'function') {
      onChange(page, pageSize);
    }
    if (typeof onPageChange === 'function') {
      onPageChange(null, { activePage: page, pageSize });
    }
  };

  return (
    <Pagination
      {...props}
      className={nextClassName}
      current={resolvedCurrent}
      total={resolvedTotalPages * 10}
      pageSize={10}
      showSizeChanger={false}
      onChange={handleChange}
    />
  );
}

export default AppPagination;
