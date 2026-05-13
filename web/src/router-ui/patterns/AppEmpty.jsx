import React from 'react';

function AppEmpty({ className = '', children }) {
  const nextClassName = ['router-empty-cell', className].filter(Boolean).join(' ');
  return <div className={nextClassName}>{children}</div>;
}

export default AppEmpty;
