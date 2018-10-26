/*==============================================================*/
/*
DataBase: AdminDB
Author:		architect.bian
Description:	AdminDB初始化数据
CreatedDate:	19:04 2013/1/10
ModifyDate:	19:04 2013/1/10
*/

/*==============================================================*/
/*
Author:		architect.bian
Description:	admin表的初始化
*/
INSERT INTO `admin` (`uid`, `adminid`, `pwd`, `dopwd`, `name`, `email`, `mobile`, `status`, `createtime`, `updatetime`) VALUES 
  ('06c41447a5284a4b977a8dddb47dcf45','xue','8d84b7abf59720d4e1b6369c99982894','8d84b7abf59720d4e1b6369c99982894','雪','1@163.com','13299999999','1','2014-05-12 18:37:42','2014-05-12 18:34:44'),
  ('e1d441056c46416098e28f012e3b3d25','zhang','43e342c4e2fe618170d011df32354163','8df3142fc535eab3ab9098bd86c04547','张','zhang@xxx.com','13683526220','1','2014-01-13 12:59:12','2014-01-13 18:27:20');

/*==============================================================*/
/*
Author:		architect.bian
Description:	adminurl表的初始化测试数据
CodeFile:		com.wbcom.web.admin.dao.***
*/
INSERT INTO `adminauthority` (`uid`, `roleuid`, `roleid`, `url`, `name`, `desc`, `createtime`, `updatetime`) VALUES 
  ('x3483b04795f48d0a70b1f9004c1f98s', 'ee483b04795f48d0a70b1f9004c1fasw', 'ROLE_ADMIN', '.*', '管理员权限', '权限描述', '2013-11-23 17:05:12', '2013-11-23 17:05:12');
