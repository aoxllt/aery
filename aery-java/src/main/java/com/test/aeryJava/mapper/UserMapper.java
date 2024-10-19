package com.test.aeryJava.mapper;

import com.test.aeryJava.pojo.UserInfo;
import org.apache.ibatis.annotations.Mapper;

@Mapper
public interface UserMapper {

    public UserInfo findUserById(Integer id);
}
