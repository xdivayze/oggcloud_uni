import { ComponentDispatchStructType } from "../../Register/Components/ComponentDispatchStruct";
import { ERR_MODE_STYLES, StatusCodes } from "../../Register/Services/utils";

export function ValidatePassword(passwordCompStruct: ComponentDispatchStructType) {
    const passwordValidStatus = ValidatePasswordImpl(passwordCompStruct)
    if (passwordValidStatus !== StatusCodes.Success ) {
        passwordCompStruct.setStyles(ERR_MODE_STYLES)
        passwordCompStruct.setText(passwordValidStatus)
        throw new Error("password invalid")
    }
    
}

function ValidatePasswordImpl(passwordCompStruct: ComponentDispatchStructType){
    if (passwordCompStruct.getRef() === null) {
        return StatusCodes.ErrNull
    }
    const password = passwordCompStruct.getRefContent().innerText
    if( /\s/.test( password)) {
        return StatusCodes.ErrWhiteSpace
    }
    if (password.length > 9) {
        return StatusCodes.ErrPasswordTooLong
    }
    return StatusCodes.Success
}