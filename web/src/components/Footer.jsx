import React, { useCallback, useEffect, useRef, useState } from 'react';
import { getFooterHTML } from '../helpers';

const Footer = () => {
  const [footer, setFooter] = useState(getFooterHTML());
  const remainCheckTimesRef = useRef(5);

  const loadFooter = useCallback(() => {
    let footer_html = localStorage.getItem('footer_html');
    if (footer_html) {
      setFooter(footer_html);
    }
  }, []);

  useEffect(() => {
    const timer = setInterval(() => {
      if (remainCheckTimesRef.current <= 0) {
        clearInterval(timer);
        return;
      }
      remainCheckTimesRef.current -= 1;
      loadFooter();
    }, 200);
    return () => clearTimeout(timer);
  }, [loadFooter]);

  // 如果未配置自定义页脚，则不渲染默认版权文案
  if (!footer) {
    return null;
  }

  return (
    <footer className='router-footer'>
      <div className='router-footer-container'>
        <div
          className='custom-footer router-footer-content'
          dangerouslySetInnerHTML={{ __html: footer }}
        ></div>
      </div>
    </footer>
  );
};

export default Footer;
