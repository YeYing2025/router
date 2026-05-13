import React from 'react';
import { Breadcrumb } from 'antd';

function AppBreadcrumb({ className = '', items = [] }) {
  const nextClassName = ['router-ui-breadcrumb', className].filter(Boolean).join(' ');
  const breadcrumbItems = items.map((item) => {
    const isLast = item?.active === true;
    const label = item?.label ?? '-';
    return {
      key: item?.key || String(label),
      title:
        item?.onClick && !isLast ? (
          <button
            type='button'
            className='router-breadcrumb-link'
            onClick={item.onClick}
          >
            {label}
          </button>
        ) : (
          <span className={`router-breadcrumb-text ${isLast ? 'active' : ''}`.trim()}>
            {label}
          </span>
        ),
    };
  });

  return <Breadcrumb className={nextClassName} separator='/' items={breadcrumbItems} />;
}

export default AppBreadcrumb;
