create function departments_code()
returns TEXT as
$$
declare
    old_id TEXT :=(select  departments_code  from departments order by departments_code desc limit 1);
    id_number char(3) :='001';
    new_id TEXT ;
    num integer;
begin
    if old_id is null then
        new_id:='dept'||id_number;
        return new_id;
    end if;
    
       num :=cast(right(old_id,3) as integer)+1;
        id_number:=
        case
            when num<10 then '00'||num
            when num<100 then '0'||num
            when num<1000 then cast(num as TEXT)
            
        end;
   
    
    
    new_id:='dept'||id_number;
    return new_id;
end; 
$$
language 'plpgsql';

