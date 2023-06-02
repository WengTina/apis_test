create function requisitions_code()
returns TEXT as
$$
declare
    old_id TEXT :=(select  requisitions_code  from requisitions order by requisitions_code desc limit 1);
    id_number char(4) :='0001';
    new_id TEXT ;
    num integer;
begin
    if old_id is null then
        new_id:='P'||id_number;
        return new_id;
    end if;
    
       num :=cast(right(old_id,4) as integer)+1;
        id_number:=
        case
            when num<10 then '000'||num
            when num<100 then '00'||num
            when num<1000 then '0'||num
            when num<10000 then cast(num as TEXT)
        end;
   
    
    
    new_id:='P'||id_number;
    return new_id;
end; 
$$
language 'plpgsql';

