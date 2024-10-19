package com.test.aeryJava.service.impl;

import com.test.aeryJava.mapper.UserMapper;
import com.test.aeryJava.pojo.UserInfo;
import com.test.aeryJava.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private UserMapper userMapper;

    @Override
    public UserInfo getUserInfo(Integer id) {
        return userMapper.findUserById(id);
    }
}
